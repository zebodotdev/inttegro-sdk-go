package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v7/otp"
	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

func TestOtpUsesTypedRequestsAndResponses(t *testing.T) {
	var requests int
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		case "/otp/initiate":
			if got := r.Header.Get("Idempotency-Key"); got != "verify-user-123" {
				t.Fatalf("expected idempotency header, got %q", got)
			}
			var params otp.InitiateParams
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				t.Fatal(err)
			}
			if params.TokenAlphabetType != otp.AlphabetTypeNumeric {
				t.Fatalf("unexpected alphabet type %q", params.TokenAlphabetType)
			}
			_, _ = w.Write([]byte(`{"transaction":{"id":"ot_1","expires_at":"2026-01-01T00:10:00Z","full_message":"Your code is {token}.","initiated_at":"2026-01-01T00:00:00Z","status":"pending_verification"}}`))
		case "/otp/verify":
			_, _ = w.Write([]byte(`{"transaction":{"id":"ot_1","expires_at":"2026-01-01T00:10:00Z","full_message":"Your code is {token}.","initiated_at":"2026-01-01T00:00:00Z","status":"verified"},"verification_attempt":{"attempted_at":"2026-01-01T00:01:00Z","id":"ov_1","presented_token":"123456","recipient":"+233241234567","result":{"verdict":"pass"}}}`))
		default:
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	transaction, err := client.Otp.Initiate(ctx, otp.InitiateParams{
		Recipient:         "+233241234567",
		ServiceName:       "MyApp",
		TokenAlphabetType: otp.AlphabetTypeNumeric,
		TokenSize:         6,
	}, request.WithIdempotencyKey("verify-user-123"))
	if err != nil {
		t.Fatal(err)
	}

	verification, err := client.Otp.Verify(ctx, otp.VerifyParams{
		Recipient:     "+233241234567",
		Token:         "123456",
		TransactionID: transaction.ID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if verification.VerificationAttempt.Result.Verdict != otp.VerificationVerdictPass {
		t.Fatalf("unexpected verdict %q", verification.VerificationAttempt.Result.Verdict)
	}
	if requests != 2 {
		t.Fatalf("expected 2 OTP requests, got %d", requests)
	}
}
