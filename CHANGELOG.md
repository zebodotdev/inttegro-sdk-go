# Changelog

## 3.0.0 - 2026-09-03

- Breaking: price creation now accepts the canonical nested `Money` amount instead of
  the obsolete flat currency and integer amount fields.
- Chime page parameters now expose the API's customer and recipient filters.
- Corrected README examples and terminology to show direct domain return values.

## 2.0.0 - 2026-09-03

- Breaking: resource methods now return domain objects and pages directly instead of response wrappers.
- Breaking: removed response-oriented exported types, including the raw HTTP response from file downloads.
- Breaking: renamed payment method and payment result types to semantic names.

## 1.0.0 - 2026-09-01

- Breaking: renamed the module and package to `github.com/zebodotdev/inttegro-sdk-go` and `inttegro`.
- Aligned documentation, examples, and the transport user agent with the public Inttegro service name.
