# Changelog

## Unreleased

## 5.0.2 - 2026-09-08

- Split large resource package files into focused files for services, request
  parameters, resources, enums, and domain-specific supporting types.
- Added source-organization checks that reject monolithic package files while
  preserving the complete v5 public API and wire behavior.

## 5.0.1 - 2026-09-08

- Grouped each resource enum's constants into one idiomatic declaration.
- Added source-shape checks that reject singleton parenthesized constant blocks
  and constants of one exported type split across multiple declarations.

## 5.0.0 - 2026-09-08

- Breaking: moved resource services, request parameters, models, and lifecycle
  values into their singular resource packages as real named Go definitions.
- Breaking: removed all root resource mirrors, deprecated compatibility aliases,
  and the plural `bankaccounts`, `paymentmethods`, and `wallets` packages.
- Breaking: changed the module path to `github.com/zebodotdev/inttegro-sdk-go/v5`.
- Added `request.Meta` and `request.WithIdempotencyKey` for cross-resource
  request controls.

## 4.5.0 - 2026-09-08

- Added singular resource packages with short domain-scoped names such as
  `product.TypeDigital`, `purchaseintent.StatusActive`,
  `refund.CreateParams`, and `order.Resource`.
- Kept the existing root resource names and client service fields source
  compatible for the remainder of v4.

## 4.4.1 - 2026-09-06

- Omit zero billing details from Order creation and validate contact fields only when billing details are supplied, matching the optional public API field.
- Decode canonical Product category, media, attributes, and price activity into the source-compatible v4 Product representation.

## 4.4.0 - 2026-09-06

- Added opt-in, typed error reporting to application-owned collectors with privacy-safe payloads, stable fingerprints, isolated reporter failures, and no reporting work when unconfigured.

## 4.3.1 - 2026-09-04

- Added the MIT license to the released module source so package documentation can be displayed by pkg.go.dev.

## 4.3.0 - 2026-09-04

- Added vendor-neutral OpenTelemetry spans for logical SDK operations, HTTP attempts, response receipt, decoding, and safe failure categories.
- Added W3C trace-context propagation plus global or per-client tracer-provider and propagator configuration.
- Kept request bodies, credentials, resource identifiers, dynamic URLs, and error details out of telemetry.

## 4.2.0 - 2026-09-03

- Added focused `wallets` and `bankaccounts` packages for financial-account variants.
- Added a `paymentmethods` package for shared mobile-money network values.
- Preserved the existing root names as type aliases for source compatibility.

## 4.1.0 - 2026-09-03

- Added the referenced product ID to returned catalog prices.

## 4.0.1 - 2026-09-03

- Corrected the order creation example to use the typed `OrderLineItemParams` request variant.

## 4.0.0 - 2026-09-03

- Breaking: replaced `OrderPaymentStatus` with the domain-level `PaymentStatus` type.
- Breaking: separated request `money.AmountParams` and `PriceParams` from returned `money.Amount`, `Price`, and `CatalogPrice` values.
- Added typed order line-item request variants and corrected purchase-intent price shapes.

## 3.0.0 - 2026-09-03

- Breaking: price creation now accepts the canonical nested `money.AmountParams` instead of
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
