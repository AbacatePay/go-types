package v2

const APIBaseURL = "https://api.abacatepay.com/"

const (
	RouteCreateCustomer = "/customers/create"
	RouteListCustomers  = "/customers/list"
	RouteGetCustomer    = "/customers/get"
	RouteDeleteCustomer = "/customers/delete"
)

const (
	RouteCreateCheckout = "/checkouts/create"
	RouteListCheckouts  = "/checkouts/list"
	RouteGetCheckout    = "/checkouts/get"
)

const (
	RouteCreatePixQRCode = "/transparents/create"
	RouteSimulatePayment = "/transparents/simulate-payment"
	RouteCheckQRCodePIX  = "/transparents/check"
)

const (
	RouteCreateCoupon       = "/coupons/create"
	RouteListCoupons        = "/coupons/list"
	RouteGetCoupon          = "/coupons/get"
	RouteDeleteCoupon       = "/coupons/delete"
	RouteToggleStatusCoupon = "/coupons/toggle"
)

const (
	RouteCreatePayout = "/payoust/create"
	RouteGetPayout    = "/payouts/get"
	RouteListPayouts  = "/payouts/list"
)

const RouteStoreDetails = "/store/get"

const (
	RouteGetMRR      = "/public-mrr/mrr"
	RouteGetMerchant = "/public-mrr/merchant-info"
	RouteGetRevenue  = "/public-mrr/revenue"
)

const (
	RouteCreateSubscription = "/subscriptions/create"
	RouteListSubscriptions  = "/subscriptions/list"
)

const (
	RouteCreateProduct = "/products/create"
	RouteListProducts  = "/products/list"
	RouteGetProduct    = "/products/get"
)
