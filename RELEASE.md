# Release Process for Reliable Proxy

This document outlines the steps to create a new release of Reliable Proxy using GoReleaser.

## Prerequisites

1. Install GoReleaser: 
   ```bash
   # macOS
   brew install goreleaser/tap/goreleaser

   # Linux
   curl -sfL https://goreleaser.com/static/get | sh
   ```

2. Ensure you have a GitHub token with `repo` scope set as an environment variable (only needed for testing):
   ```bash
   export GITHUB_TOKEN=your_github_token
   ```

## Creating a Release

1. Make sure your code is ready for release:
   - All changes are committed and pushed
   - All tests pass
   - Documentation is up to date

2. Tag the new release:
   ```bash
   # For a new version (e.g., v1.0.0)
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

3. The GitHub Actions workflow will automatically:
   - Build binaries for all supported platforms
   - Create a GitHub release with the binaries
   - Generate a changelog based on commits

4. Verify the release on the GitHub Releases page.

## Testing a Release Locally

To test the release process locally without publishing:

```bash
goreleaser release --snapshot --clean --skip=publish
```

This will build the binaries and create a local snapshot of the release in the `dist` directory.

## Using the Homebrew Tap (Optional)

To enable Homebrew distribution:

1. Create a repository for your Homebrew tap (e.g., `homebrew-tap`)
2. Uncomment and update the Homebrew configuration in both:
   - `.goreleaser.yaml`
   - `.github/workflows/release.yml`
3. Set up the required secrets in your GitHub repository settings

## Troubleshooting

- If the release fails, check the GitHub Actions logs for errors
- For local testing issues, run with verbose output: `goreleaser release --snapshot --clean --skip=publish -v` 