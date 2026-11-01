# Vendored Swagger UI

These are the static assets of [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist), vendored here
because `github.com/swaggo/files/v2` (which Sherlock used before) is stuck on Swagger UI 5.18.2 and has no
newer release.

- Version: **5.34.1** (bundles DOMPurify 3.4.17)
- Source: `https://registry.npmjs.org/swagger-ui-dist/-/swagger-ui-dist-5.34.1.tgz`
- npm integrity: `sha512-VSr8fXfI93obYaQicLJbLUKGRxDojj/SfRJDDLCioGTUBj4BK/Nj4eE9jCeCAnJxVLi7iS+08PbiZLyrQKBkJQ==`
- Provenance: verified with `npm audit signatures` (registry signatures and the SLSA provenance attestation). The
  attestation's signed sha512 digest matches this tarball, and it records the build as a GitHub-hosted Actions run
  of `.github/workflows/release-swagger-ui-dist.yml` in `swagger-api/swagger-ui` at commit
  `d3296954a6b813898d298815ac3e68daa1cf7408`. The files were not rebuilt from source.

Only the files that Sherlock serves are vendored. `swaggo/gin-swagger` generates `index.html`, `index.css`,
`swagger-initializer.js` and `doc.json` itself, so those are not here. The source maps and the ES bundles are
left out because nothing loads them.

## Updating

Nothing tracks this directory automatically (dependabot cannot see it), so check it when Swagger UI or DOMPurify
publish security advisories.

```bash
V=<new version>
curl -sO "https://registry.npmjs.org/swagger-ui-dist/-/swagger-ui-dist-$V.tgz"
# compare against the registry before trusting it:
shasum "swagger-ui-dist-$V.tgz"; curl -s "https://registry.npmjs.org/swagger-ui-dist/$V" | grep '"shasum"'
tar xzf "swagger-ui-dist-$V.tgz"
# then copy the same files that are in dist/ from package/ into dist/, and update the version above
```

After updating, load `/swagger/index.html` in a browser and check the console for Content Security Policy errors
(see `internal/middleware/security`).
