# TikTokMall
#### ByteDance Youth Training Camp Douyin E-commerce Project (One of the Highest-Scoring Backend Projects)



backend: Backend

The script requires `air` to be installed for startup
Startup script:
```shell
cd backend
make start
```


frontend: Frontend

node version v14.17.5

```shell
cd frontend
npm install
npm run serve
```



## I. Project Overview
This project is dedicated to building a feature-rich, AI-powered smart e-commerce shopping platform with standardized development and deployment processes. It covers functional modules including user authentication, user management, product management, shopping cart, order processing, payment settlement (supporting bank cards and Alipay), and LLM intelligent customer service, aiming to provide users with a seamless shopping experience and intelligent services.

Apifox API Documentation: wf1rbhp3a3.apifox.cn

Documentation Address: [Project Description Document](https://rnutm2ali2.feishu.cn/docx/ZHXBdl62Doeql1xLNQ2conUWnTf)

![Overall Architecture Diagram](https://hrs-aliyunoss.oss-cn-beijing.aliyuncs.com/imges/image-20250305211401954.png)

## II. Module Completion Status

### (I) Authentication Center
* [x]  Issue identity tokens
* [x]  Validate identity tokens
* [ ]  Renew identity tokens (Advanced)

### (II) User Service
* [x]  Create user: Completed
* [x]  Login: Completed
* [x]  User logout (Optional)
* [x]  Delete user (Optional)
* [x]  Update user (Optional)
* [x]  Get user identity information

### (III) Product Service
* [x]  Product search
* [x]  Product categorization
* [x]  Merchant feature: Create product (Optional)
* [x]  Merchant feature: Modify product information (Optional)
* [x]  Merchant feature: Delete product (Optional)
* [x]  Query product information (single product, batch products)
* [ ]  Product detail video carousel display (Optional)

### (IV) Shopping Cart Service
* [x]  Create shopping cart
* [x]  Clear shopping cart
* [x]  Get shopping cart information
* [x]  Shopping cart product quantity modification interface
* [x]  Shopping cart product deletion interface

### (V) Order Service
* [x]  Create order
* [x]  Query order list
* [ ]  Modify order information (Optional)
* [x]  Order scheduled cancellation (Advanced)

### (VI) Settlement
* [x]  Order settlement
* [x]  Awaiting payment
* [x]  Payment result query

### (VII) Payment
* [x]  Cancel payment (Advanced)
* [x]  Process payment
* [x]  Scheduled payment cancellation (Advanced)

### (VIII) AI Large Model
* [x]  Order query (Optional)
* [x]  Simulated auto-ordering (Optional)
* [x]  Product recommendation (Optional)

### (IX) Frontend
* [x]  API integration
* [x]  Merchant features (Optional)
* [x]  User information page
* [ ]  Product video display page (Optional)


## III. Project Progress and Achievements
Currently, the overall project progress is smooth, and most core features have been developed and tested, making it ready for deployment and operation. Through the collaboration of various service modules, the platform is able to provide users with a complete shopping process—from user authentication, product browsing, shopping cart management, order creation and settlement, to payment completion, as well as AI large model-assisted features—enhancing user experience and the platform's intelligent level.
