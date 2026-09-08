package payout

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

// PayoutsService manages payout configuration, scheduling, and history.
//
// Payouts move funds from your Inttegro balance to your connected financial
// accounts (mobile money, bank, Dosh). Use this service to:
//
//   - Configure payout destination accounts
//   - View payout settings and schedule
//   - Switch between automatic and manual payout modes
//   - Enable/disable currency conversion
//   - List payout history
//
// Example:
//
//	// Configure payout destinations
//	settings, err := client.Payouts.SetDestinations(ctx, map[string]string{
//	    "ghs": "fa_abc123",  // GHS payouts go to this mobile money account
//	    "usd": "fa_def456",  // USD payouts go to this bank account
//	})
//
// Learn more: https://studio.inttegro.com/set-up-payouts
type Service struct {
	client transport.Client
}

// SetDestinations configures which financial accounts receive payouts by currency.
//
// Map each currency you accept to a financial account ID. When balance
// transactions in that currency become eligible, they're paid out to the
// corresponding account.
//
// Parameters:
//   - destinations: Map of currency code to financial account ID
//     Example: {"ghs": "fa_abc123", "usd": "fa_def456"}
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.SetDestinations(ctx, map[string]string{
//	    "ghs": "fa_abc123",
//	})
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("GHS destination: %s\n", settings.Destinations["ghs"])
//
// Learn more: https://studio.inttegro.com/set-payout-destinations
func (s *Service) SetDestinations(ctx context.Context, destinations map[string]string) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	payload := struct {
		Destinations map[string]string `json:"destinations"`
	}{Destinations: destinations}
	if err := s.client.Do(ctx, "POST", "/payouts/set_destinations", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Settings retrieves your current payout configuration.
//
// Returns payout schedule, destination accounts, and FX settings.
//
// Example:
//
//	settings, err := client.Payouts.Settings(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Schedule: %s\n", settings.Schedule.Type)
//	fmt.Printf("Destinations: %v\n", settings.Destinations)
func (s *Service) Settings(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/settings", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Schedule creates a payout to a connected financial account.
func (s *Service) Schedule(ctx context.Context, params ScheduleParams) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/schedule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// Lookup retrieves a payout by ID.
func (s *Service) Lookup(ctx context.Context, payoutID string) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/lookup", map[string]string{"payout_id": payoutID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// DisableAutomatic switches to manual payout mode.
//
// In manual mode, payouts only happen when you explicitly request them.
// Use this for marketplace platforms or when you need precise control
// over payout timing.
//
// Returns the updated payout settings with manual schedule.
//
// Example:
//
//	settings, err := client.Payouts.DisableAutomatic(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Schedule type: %s\n", settings.Schedule.Type) // "manual"
//
// Learn more: https://studio.inttegro.com/disable-automatic-payouts
func (s *Service) DisableAutomatic(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/disable", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// EnableAutomatic switches payout scheduling back to automatic mode.
func (s *Service) EnableAutomatic(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/enable", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// EnableFX enables currency conversion for payouts.
//
// When enabled, you can receive payouts in a different currency than
// your source funds. Requires FX-enabled destination accounts.
//
// Example: Accept USD payments, receive GHS payouts.
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.EnableFX(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("FX enabled: %v\n", settings.FxEnabled) // true
func (s *Service) EnableFX(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/enable_fx", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// DisableFX disables currency conversion for payouts.
//
// When disabled, you can only receive payouts in the same currency as
// your source funds.
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.DisableFX(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("FX enabled: %v\n", settings.FxEnabled) // false
func (s *Service) DisableFX(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/disable_fx", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Cancel cancels a scheduled payout before execution.
//
// Only payouts in "scheduled" status with a future execution time can be
// canceled. Once canceled, included balance transactions remain available for
// future payouts.
//
// Parameters:
//   - payoutID: Scheduled payout ID to cancel
//
// Returns the canceled payout object.
func (s *Service) Cancel(ctx context.Context, payoutID string) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	payload := struct {
		PayoutID string `json:"payout_id"`
	}{PayoutID: payoutID}
	if err := s.client.Do(ctx, "POST", "/payouts/cancel", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// Page returns a paginated list of recent payouts.
//
// View payout history including amounts, statuses, and timing. Results
// are sorted by initiation date (newest first).
//
// Parameters:
//   - params.PageNumber: Page to retrieve (optional, default: 1)
//   - params.PageSize: Payouts per page (optional, default: 20, max: 100)
//
// Returns a slice of payouts for the requested page.
//
// Example:
//
//	payouts, err := client.Payouts.Page(ctx, payout.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	})
//	for _, payout := range payouts {
//	    fmt.Printf("Payout %s: %s %d %s\n",
//	        payout.ID, payout.Amount.Currency,
//	        payout.Amount.Value, payout.Status)
//	}
func (s *Service) Page(ctx context.Context, params PageParams) ([]Resource, error) {
	var resp struct {
		Page struct {
			Payouts []Resource `json:"payouts"`
		} `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Payouts, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
