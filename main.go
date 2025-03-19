package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type RegionInfo struct {
	Region      string `json:"region"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country_name"`
	City        string `json:"city"`
}

var (
	cachedRegion     string
	regionCacheFile  string
	regionCacheMutex sync.RWMutex
)

func init() {
	if _, err := os.Stat("/app/data"); err == nil {
		regionCacheFile = "/app/data/region_cache.txt"
	} else {
		regionCacheFile = "region_cache.txt"
	}
}

func loadEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
}

func fetchAndCacheRegion() {
	if region, err := ioutil.ReadFile(regionCacheFile); err == nil && len(region) > 0 {
		regionCacheMutex.Lock()
		cachedRegion = string(region)
		regionCacheMutex.Unlock()
		return
	}

	region := getRegion()
	regionCacheMutex.Lock()
	cachedRegion = region
	regionCacheMutex.Unlock()
	
	if region != "" {
		ioutil.WriteFile(regionCacheFile, []byte(region), 0644)
	}
}

func getRegion() string {
	geoServices := []string{"https://ipapi.co/json/", "https://ipinfo.io/json"}
	results := make(chan string, len(geoServices))
	var wg sync.WaitGroup
	client := &http.Client{Timeout: 3 * time.Second}
	
	for _, service := range geoServices {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			
			resp, err := client.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				if err == nil {
					resp.Body.Close()
				}
				return
			}
			defer resp.Body.Close()
			
			var info RegionInfo
			if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
				return
			}
			
			regionParts := []string{}
			if info.City != "" {
				regionParts = append(regionParts, info.City)
			}
			if info.Region != "" {
				regionParts = append(regionParts, info.Region)
			}
			if info.Country != "" {
				regionParts = append(regionParts, info.Country)
			} else if info.CountryCode != "" {
				regionParts = append(regionParts, info.CountryCode)
			}
			
			if len(regionParts) > 0 {
				results <- strings.Join(regionParts, ", ")
			}
		}(service)
	}
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	select {
	case result := <-results:
		return result
	case <-time.After(3100 * time.Millisecond):
		return ""
	}
}

func respondWithError(w http.ResponseWriter, msg string, code int) {
	http.Error(w, fmt.Sprintf(`{"error":"%s"}`, msg), code)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetAPIURL := r.Header.Get("target-api-url")
	if targetAPIURL == "" {
		targetAPIURL = os.Getenv("TARGET_API_URL")
	}
	
	if r.URL.Path == "/" && targetAPIURL == "" {
		w.Header().Set("Content-Type", "application/json")
		
		regionCacheMutex.RLock()
		region := cachedRegion
		regionCacheMutex.RUnlock()
		
		statusMsg := `{"status":"Reliable Proxy server is running"}`
		if region != "" {
			statusMsg = fmt.Sprintf(`{"status":"Reliable Proxy server is running","region":"%s"}`, region)
		}
		
		w.Write([]byte(statusMsg))
		return
	}
	
	if targetAPIURL == "" {
		respondWithError(w, "Missing target-api-url header or TARGET_API_URL environment variable", http.StatusBadRequest)
		return
	}

	targetURL := strings.TrimRight(targetAPIURL, "/") + "/" + strings.TrimLeft(r.URL.Path, "/")
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		respondWithError(w, "Invalid URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	parsedURL.RawQuery = r.URL.RawQuery

	proxyReq, err := http.NewRequest(r.Method, parsedURL.String(), r.Body)
	if err != nil {
		respondWithError(w, "Error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	excludedHeaders := map[string]bool{"host": true, "target-api-url": true}
	for name, values := range r.Header {
		if !excludedHeaders[strings.ToLower(name)] {
			proxyReq.Header[name] = values
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(proxyReq)
	if err != nil {
		respondWithError(w, "Error sending request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	loadEnvFile()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Reliable Proxy running on :%s", port)
	
	go func() {
		fetchAndCacheRegion()
		regionCacheMutex.RLock()
		region := cachedRegion
		regionCacheMutex.RUnlock()
		
		if region != "" {
			log.Printf("Proxy region: %s", region)
		}
	}()
	
	http.HandleFunc("/", proxyHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server error: %s", err)
	}
} 
