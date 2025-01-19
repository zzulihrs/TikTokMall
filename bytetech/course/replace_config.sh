cd app/cart
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/checkout
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/email
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/frontend
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/order
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/payment
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/product
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..
cd app/user
# go mod edit --replace bytetech/course/common/=../../common
go mod tidy
cd ../..