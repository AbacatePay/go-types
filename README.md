<h1 align="center">AbacatePay Go Types</h1>

<p align="center">
Types, structs e helpers oficiais para integração com a API da AbacatePay em Go.
</p>

## 📦 Instalação

```bash
go get github.com/almeidazs/go-abacate-types@latest
```

<h2 align="center">Como a AbacatePay API Types documenta</h2>

Antes de tudo, você deve específicar a versão da API que você deseja importar os tipos. Coloque `/v*` no final da importação, sendo `*` a versão que deseja usar:

```go
import "github.com/almeidazs/go-abacate-types/v1"
```

- Prefixo `API*`
Representa estruturas gerais da API (Objetos retornados, modelos internos etc.).

- Prefixo `Webhook*`
Representa payloads recebidos pelos eventos de webhook.
Documentação: https://docs.abacatepay.com/pages/webhooks

- Prefixo `REST<HTTPMethod>*`
Tipos usados em requisições diretas à API.
  - Sufixo Body → corpo enviado na requisição
  Ex.: `RESTPostCreateNewChargeBody`

  - Sufixo `QueryParams` → parâmetros de query
  Ex.: `RESTGetCheckQRCodePixStatusQueryParams`

  - Sufixo `Data` → dados retornados pela API
  Ex.: `RESTGetListCouponsData`

- O pacote **NÃO adiciona tipos além do que existe na documentação oficial**.
Cada tipo reflete exatamente o que está documentado aqui:
https://docs.abacatepay.com/pages/introduction

- Campos marcados com `@unstable`
São campos que não têm definição formal na documentação, mas cujo tipo foi inferido com base nos exemplos oficiais.
(Ex.: `WebhookWithdrawDoneEvent.billing.kind`)

<h2 align="center">Quickstart</h2>

<p align="center"><strong>Crie um novo cupom</strong></p>

```go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/almeidazs/go-abacate-types/v1"
)

func CreateCoupon(body v1.RESTPostCreateCouponBody) (*v1.APICoupon, error) {
	url := v1.APIBaseURL + v1.RouteCreateCoupon

	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result v1.RESTPostCreateCouponData

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
```

<p align="center"><strong>Receber Webhooks da AbacatePay</strong></p>

```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/almeidazs/go-abacate-types/v1"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var event v1.WebhookEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	switch event.Event {
	case v1.WebhookEventBillingPaid:
		// Pagamento confirmado
	case v1.WebhookEventWithdrawDone:
		// Saque concluído
	case v1.WebhookEventWithdrawFailed:
		// Saque falhou
	}

	w.WriteHeader(http.StatusOK)
}
```
