## Stripe POC

### Authentication

- Customer creates restricted API key  
  - we use key to get historical transactions
- We provide Customer with SDK
  - he adds his client id and secret in the SDK
  - it creates a webhook with our URL to provides us with any new transactions