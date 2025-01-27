resource "aws_sns_topic" "order_created_topic" {
  name = "order-created"
}

resource "aws_sns_topic" "order_updated_topic" {
  name = "order-updated"
}
