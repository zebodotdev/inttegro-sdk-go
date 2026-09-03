package inttegro

// This file exposes every string enum in the public API as a typed constant.
// Use these symbols in requests instead of spelling wire values by hand.

type PaymentMethodType string

const (
	PaymentMethodTypeMobileMoney PaymentMethodType = "mobile_money"
	PaymentMethodTypeBankAccount PaymentMethodType = "bank_account"
	PaymentMethodTypeCard        PaymentMethodType = "card"
	PaymentMethodTypeMotito      PaymentMethodType = "motito"
)

type LineItemType string

const (
	LineItemTypeProduct  LineItemType = "product"
	LineItemTypeFee      LineItemType = "fee"
	LineItemTypeShipping LineItemType = "shipping"
)

type ChimeRecipientType string

const (
	ChimeRecipientTypePhone ChimeRecipientType = "phone"
	ChimeRecipientTypeEmail ChimeRecipientType = "email"
)

type ChimeTransport string

const (
	ChimeTransportSMS   ChimeTransport = "sms"
	ChimeTransportEmail ChimeTransport = "email"
)

type FinancialAccountType string

const (
	FinancialAccountTypeWallet FinancialAccountType = "wallet"
	FinancialAccountTypeBank   FinancialAccountType = "bank_account"
	FinancialAccountTypeDosh   FinancialAccountType = "dosh_account"
)

type BalanceTransactionType string

const (
	BalanceTransactionTypePayment BalanceTransactionType = "payment"
	BalanceTransactionTypeRefund  BalanceTransactionType = "refund"
)

type AppManagementRole string

const (
	AppManagementRoleParent AppManagementRole = "parent"
	AppManagementRoleChild  AppManagementRole = "child"
)

type AppCredentialOwner string

const (
	AppCredentialOwnerChild  AppCredentialOwner = "child"
	AppCredentialOwnerParent AppCredentialOwner = "parent"
)

type AppRelationshipKind string

const AppRelationshipKindPlacement AppRelationshipKind = "placement"

type AppRelationshipStatus string

const (
	AppRelationshipStatusActive    AppRelationshipStatus = "active"
	AppRelationshipStatusInactive  AppRelationshipStatus = "inactive"
	AppRelationshipStatusSuspended AppRelationshipStatus = "suspended"
	AppRelationshipStatusRevoked   AppRelationshipStatus = "revoked"
)

type SecretKeyTokenType string

const SecretKeyTokenTypeBearer SecretKeyTokenType = "bearer"

type SecretKeyStatus string

const (
	SecretKeyStatusActive  SecretKeyStatus = "active"
	SecretKeyStatusRevoked SecretKeyStatus = "revoked"
	SecretKeyStatusExpired SecretKeyStatus = "expired"
)

type SecretKeyAuthResult string

const (
	SecretKeyAuthResultSucceeded SecretKeyAuthResult = "succeeded"
	SecretKeyAuthResultFailed    SecretKeyAuthResult = "failed"
)

type FileStatus string

const (
	FileStatusUploading  FileStatus = "uploading"
	FileStatusProcessing FileStatus = "processing"
	FileStatusAvailable  FileStatus = "available"
	FileStatusFailed     FileStatus = "failed"
	FileStatusDeleted    FileStatus = "deleted"
)

type FileDisposition string

const (
	FileDispositionAttachment FileDisposition = "attachment"
	FileDispositionInline     FileDisposition = "inline"
)

type FileDelivery string

const (
	FileDeliveryStream   FileDelivery = "stream"
	FileDeliveryRedirect FileDelivery = "redirect"
)

type FileScanStatus string

const (
	FileScanStatusPending FileScanStatus = "pending"
	FileScanStatusPassed  FileScanStatus = "passed"
	FileScanStatusFailed  FileScanStatus = "failed"
	FileScanStatusSkipped FileScanStatus = "skipped"
)

type FileSourceType string

const (
	FileSourceTypeDirect        FileSourceType = "direct"
	FileSourceTypeUploadRequest FileSourceType = "upload_request"
	FileSourceTypeService       FileSourceType = "service"
)

type FileStorageEncoding string

const (
	FileStorageEncodingIdentity FileStorageEncoding = "identity"
	FileStorageEncodingBrotli   FileStorageEncoding = "br"
)

type FileLinkStatus string

const (
	FileLinkStatusActive   FileLinkStatus = "active"
	FileLinkStatusRevoked  FileLinkStatus = "revoked"
	FileLinkStatusExpired  FileLinkStatus = "expired"
	FileLinkStatusDisabled FileLinkStatus = "disabled"
)

type FileLinkKind string

const FileLinkKindPublic FileLinkKind = "public"

type FileLinkDeliveryMode string

const (
	FileLinkDeliveryModeRedirect FileLinkDeliveryMode = "redirect"
	FileLinkDeliveryModeDownload FileLinkDeliveryMode = "download"
	FileLinkDeliveryModeInline   FileLinkDeliveryMode = "inline"
)

type UploadRequestStatus string

const (
	UploadRequestStatusPending   UploadRequestStatus = "pending"
	UploadRequestStatusUploading UploadRequestStatus = "uploading"
	UploadRequestStatusFulfilled UploadRequestStatus = "fulfilled"
	UploadRequestStatusExpired   UploadRequestStatus = "expired"
	UploadRequestStatusCanceled  UploadRequestStatus = "canceled"
	UploadRequestStatusFailed    UploadRequestStatus = "failed"
)

type UploadReviewDecision string

const (
	UploadReviewDecisionApproved UploadReviewDecision = "approved"
	UploadReviewDecisionRejected UploadReviewDecision = "rejected"
)

type UploadReviewType string

const (
	UploadReviewTypeAutomatic UploadReviewType = "automatic"
	UploadReviewTypeManual    UploadReviewType = "manual"
)

type PaymentNextActionType string

const (
	PaymentNextActionTypeConfirmPayment PaymentNextActionType = "confirm_payment"
	PaymentNextActionTypeExecute        PaymentNextActionType = "execute"
	PaymentNextActionTypeRedirect       PaymentNextActionType = "redirect"
	PaymentNextActionTypeAuthorize      PaymentNextActionType = "authorize"
	PaymentNextActionTypeNone           PaymentNextActionType = "none"
)

type PaymentConfirmationChannel string

const (
	PaymentConfirmationChannelSMS   PaymentConfirmationChannel = "sms"
	PaymentConfirmationChannelEmail PaymentConfirmationChannel = "email"
	PaymentConfirmationChannelPush  PaymentConfirmationChannel = "push"
)

type MobileMoneyNetwork string

const (
	MobileMoneyNetworkAirtel   MobileMoneyNetwork = "airtel"
	MobileMoneyNetworkMTN      MobileMoneyNetwork = "mtn"
	MobileMoneyNetworkTelecel  MobileMoneyNetwork = "telecel"
	MobileMoneyNetworkVodafone MobileMoneyNetwork = "vodafone"
)

type ProductType string

const (
	ProductTypePhysical ProductType = "physical"
	ProductTypeDigital  ProductType = "digital"
	ProductTypeService  ProductType = "service"
	ProductTypeVoucher  ProductType = "voucher"
	ProductTypeCustom   ProductType = "custom"
	ProductTypeCause    ProductType = "cause"
)

type ProductShipmentType string

const (
	ProductShipmentTypeDelivery ProductShipmentType = "delivery"
	ProductShipmentTypeDownload ProductShipmentType = "download"
	ProductShipmentTypeRender   ProductShipmentType = "render"
	ProductShipmentTypeService  ProductShipmentType = "service"
	ProductShipmentTypeStream   ProductShipmentType = "stream"
)

type ProductShipmentInputType string

const (
	ProductShipmentInputTypeDelivery ProductShipmentInputType = "delivery"
	ProductShipmentInputTypeDownload ProductShipmentInputType = "download"
	ProductShipmentInputTypeRender   ProductShipmentInputType = "render"
	ProductShipmentInputTypeStream   ProductShipmentInputType = "stream"
)

type PurchaseIntentStatus string

const (
	PurchaseIntentStatusActive   PurchaseIntentStatus = "active"
	PurchaseIntentStatusExpired  PurchaseIntentStatus = "expired"
	PurchaseIntentStatusInactive PurchaseIntentStatus = "inactive"
	PurchaseIntentStatusUsed     PurchaseIntentStatus = "used"
)

type PurchaseIntentActivityType string

const (
	PurchaseIntentActivityTypeExpiredViewed  PurchaseIntentActivityType = "expired_viewed"
	PurchaseIntentActivityTypeOrderCreated   PurchaseIntentActivityType = "order_created"
	PurchaseIntentActivityTypePaymentFailed  PurchaseIntentActivityType = "payment_failed"
	PurchaseIntentActivityTypePaymentStarted PurchaseIntentActivityType = "payment_started"
	PurchaseIntentActivityTypeViewed         PurchaseIntentActivityType = "viewed"
)

type WalletType string

const WalletTypeMobileMoney WalletType = "mobile_money"

type BankAccountType string

const BankAccountTypeGhanaBankAccount BankAccountType = "ghana_bank_account"

type MessageTemplateChannel string

const (
	MessageTemplateChannelSMS   MessageTemplateChannel = "sms"
	MessageTemplateChannelEmail MessageTemplateChannel = "email"
)

type MessageTemplateStatus string

const (
	MessageTemplateStatusDraft     MessageTemplateStatus = "draft"
	MessageTemplateStatusPublished MessageTemplateStatus = "published"
	MessageTemplateStatusArchived  MessageTemplateStatus = "archived"
)

type MessageTemplateVariableType string

const (
	MessageTemplateVariableTypeString   MessageTemplateVariableType = "string"
	MessageTemplateVariableTypeNumber   MessageTemplateVariableType = "number"
	MessageTemplateVariableTypeInteger  MessageTemplateVariableType = "integer"
	MessageTemplateVariableTypeBoolean  MessageTemplateVariableType = "boolean"
	MessageTemplateVariableTypeURL      MessageTemplateVariableType = "url"
	MessageTemplateVariableTypeEmail    MessageTemplateVariableType = "email"
	MessageTemplateVariableTypePhone    MessageTemplateVariableType = "phone"
	MessageTemplateVariableTypeDate     MessageTemplateVariableType = "date"
	MessageTemplateVariableTypeDatetime MessageTemplateVariableType = "datetime"
	MessageTemplateVariableTypeArray    MessageTemplateVariableType = "array"
)

type MessageTemplateVariableItemType string

const (
	MessageTemplateVariableItemTypeString   MessageTemplateVariableItemType = "string"
	MessageTemplateVariableItemTypeNumber   MessageTemplateVariableItemType = "number"
	MessageTemplateVariableItemTypeInteger  MessageTemplateVariableItemType = "integer"
	MessageTemplateVariableItemTypeBoolean  MessageTemplateVariableItemType = "boolean"
	MessageTemplateVariableItemTypeURL      MessageTemplateVariableItemType = "url"
	MessageTemplateVariableItemTypeEmail    MessageTemplateVariableItemType = "email"
	MessageTemplateVariableItemTypePhone    MessageTemplateVariableItemType = "phone"
	MessageTemplateVariableItemTypeDate     MessageTemplateVariableItemType = "date"
	MessageTemplateVariableItemTypeDatetime MessageTemplateVariableItemType = "datetime"
)

type ContentSafetyStatus string

const (
	ContentSafetyStatusAllowed     ContentSafetyStatus = "allowed"
	ContentSafetyStatusRejected    ContentSafetyStatus = "rejected"
	ContentSafetyStatusQuarantined ContentSafetyStatus = "quarantined"
)

type OrderDocumentKind string

const (
	OrderDocumentKindInvoice OrderDocumentKind = "invoice"
	OrderDocumentKindReceipt OrderDocumentKind = "receipt"
)

type DeliveryChannel string

const (
	DeliveryChannelEmail DeliveryChannel = "email"
	DeliveryChannelSMS   DeliveryChannel = "sms"
)

type CheckoutOrderStatus string

const (
	CheckoutOrderStatusPreparing       CheckoutOrderStatus = "preparing"
	CheckoutOrderStatusRequiresPayment CheckoutOrderStatus = "requires_payment"
	CheckoutOrderStatusCompleted       CheckoutOrderStatus = "completed"
	CheckoutOrderStatusCanceled        CheckoutOrderStatus = "canceled"
	CheckoutOrderStatusExpired         CheckoutOrderStatus = "expired"
)

type OrderStatus string

const (
	OrderStatusPreparing       OrderStatus = "preparing"
	OrderStatusRequiresPayment OrderStatus = "requires_payment"
	OrderStatusPaid            OrderStatus = "paid"
	OrderStatusCompleted       OrderStatus = "completed"
	OrderStatusCanceled        OrderStatus = "canceled"
	OrderStatusExpired         OrderStatus = "expired"
	OrderStatusUnknown         OrderStatus = "unknown"
)

type PaymentStatus string

const (
	PaymentStatusInitiated      PaymentStatus = "initiated"
	PaymentStatusRequiresAction PaymentStatus = "requires_action"
	PaymentStatusOverdue        PaymentStatus = "overdue"
	PaymentStatusExecuted       PaymentStatus = "executed"
	PaymentStatusPaid           PaymentStatus = "paid"
	PaymentStatusCanceled       PaymentStatus = "canceled"
	PaymentStatusExpired        PaymentStatus = "expired"
	PaymentStatusFailed         PaymentStatus = "failed"
	PaymentStatusUnknown        PaymentStatus = "unknown"
)

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusInitiated PaymentAttemptStatus = "initiated"
	PaymentAttemptStatusExecuted  PaymentAttemptStatus = "executed"
	PaymentAttemptStatusSucceeded PaymentAttemptStatus = "succeeded"
	PaymentAttemptStatusCanceled  PaymentAttemptStatus = "canceled"
	PaymentAttemptStatusExpired   PaymentAttemptStatus = "expired"
	PaymentAttemptStatusFailed    PaymentAttemptStatus = "failed"
	PaymentAttemptStatusUnknown   PaymentAttemptStatus = "unknown"
)

type CheckoutPaymentStatus string

const (
	CheckoutPaymentStatusRequiresAction CheckoutPaymentStatus = "requires_action"
	CheckoutPaymentStatusProcessing     CheckoutPaymentStatus = "processing"
	CheckoutPaymentStatusSucceeded      CheckoutPaymentStatus = "succeeded"
	CheckoutPaymentStatusFailed         CheckoutPaymentStatus = "failed"
	CheckoutPaymentStatusCancelled      CheckoutPaymentStatus = "cancelled"
)

type PaymentResultStatus string

const (
	PaymentResultStatusPending              PaymentResultStatus = "pending"
	PaymentResultStatusRequiresConfirmation PaymentResultStatus = "requires_confirmation"
	PaymentResultStatusProcessing           PaymentResultStatus = "processing"
	PaymentResultStatusSucceeded            PaymentResultStatus = "succeeded"
	PaymentResultStatusFailed               PaymentResultStatus = "failed"
)

type OrderCreatedFromResourceType string

const OrderCreatedFromResourceTypePurchaseIntent OrderCreatedFromResourceType = "purchase_intent"

type RefundReason string

const (
	RefundReasonRequestedByCustomer RefundReason = "requested_by_customer"
	RefundReasonDuplicate           RefundReason = "duplicate"
	RefundReasonFraudulent          RefundReason = "fraudulent"
	RefundReasonOrderCanceled       RefundReason = "order_canceled"
	RefundReasonItemReturned        RefundReason = "item_returned"
	RefundReasonItemDamaged         RefundReason = "item_damaged"
	RefundReasonItemNotReceived     RefundReason = "item_not_received"
	RefundReasonItemNotAsDescribed  RefundReason = "item_not_as_described"
	RefundReasonCustom              RefundReason = "custom"
)

type RefundStatus string

const (
	RefundStatusCanceled   RefundStatus = "canceled"
	RefundStatusFailed     RefundStatus = "failed"
	RefundStatusPending    RefundStatus = "pending"
	RefundStatusProcessing RefundStatus = "processing"
	RefundStatusSucceeded  RefundStatus = "succeeded"
)

type PayoutStatus string

const (
	PayoutStatusInitialized PayoutStatus = "initialized"
	PayoutStatusScheduled   PayoutStatus = "scheduled"
	PayoutStatusProcessing  PayoutStatus = "processing"
	PayoutStatusExecuting   PayoutStatus = "executing"
	PayoutStatusSucceeded   PayoutStatus = "succeeded"
	PayoutStatusInvalid     PayoutStatus = "invalid"
	PayoutStatusCanceled    PayoutStatus = "canceled"
)

type ChimeEmailSchemaKind string

const (
	ChimeEmailSchemaKindGmailViewAction  ChimeEmailSchemaKind = "gmail_view_action"
	ChimeEmailSchemaKindSchemaOrgOrder   ChimeEmailSchemaKind = "schema_org_order"
	ChimeEmailSchemaKindSchemaOrgInvoice ChimeEmailSchemaKind = "schema_org_invoice"
)

type OTPAlphabetType string

const (
	OTPAlphabetTypeNumeric      OTPAlphabetType = "numeric"
	OTPAlphabetTypeAlpha        OTPAlphabetType = "alpha"
	OTPAlphabetTypeAlphanumeric OTPAlphabetType = "alphanumeric"
)

type OTPStatus string

const (
	OTPStatusCanceled            OTPStatus = "canceled"
	OTPStatusExpired             OTPStatus = "expired"
	OTPStatusPending             OTPStatus = "pending"
	OTPStatusPendingDelivery     OTPStatus = "pending_delivery"
	OTPStatusPendingVerification OTPStatus = "pending_verification"
	OTPStatusVerified            OTPStatus = "verified"
)

type OTPTransmissionStatus string

const (
	OTPTransmissionStatusDelivered OTPTransmissionStatus = "delivered"
	OTPTransmissionStatusFailed    OTPTransmissionStatus = "failed"
	OTPTransmissionStatusSubmitted OTPTransmissionStatus = "submitted"
)

type OTPVerificationVerdict string

const (
	OTPVerificationVerdictFail OTPVerificationVerdict = "fail"
	OTPVerificationVerdictPass OTPVerificationVerdict = "pass"
)
