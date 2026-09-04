# Changelog

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
