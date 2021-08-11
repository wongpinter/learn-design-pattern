package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatalf("A payment method type 'Cash' was not found: %v", err)
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment message type 'Cash' wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatalf("A payment method type 'DebitCard' was not found: %v", err)
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Errorf("The cash payment message type 'DebitCard' wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestPaymentMethodNotExists(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Errorf("A payment method with ID 20 most return error")
	}

	t.Log("LOG:", err)
}
