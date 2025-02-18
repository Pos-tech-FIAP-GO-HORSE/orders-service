package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/async/contract"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/product_service"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type Handler struct {
	orderService   order_service.IOrderService
	productService product_service.IProductService
}

func NewHandler(orderService order_service.IOrderService, productService product_service.IProductService) IHandler {
	return &Handler{orderService, productService}
}

func (h *Handler) Handle(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		sns := record.SNS

		switch {
		case strings.Contains(sns.TopicArn, "payments-service-events-payment-created"):
			var paymentEvent contract.PaymentEvent
			if err := json.Unmarshal([]byte(sns.Message), &paymentEvent); err != nil {
				return err
			}

			order, err := h.orderService.FindByPublicID(ctx, paymentEvent.PublicID)
			if err != nil {
				return err
			}

			orderUpdated, err := h.orderService.UpdateStatusByID(ctx, order.ID, paymentEvent.Status)
			if err != nil {
				return err
			}

			zap.L().Info("order status updated successfully", zap.Any("order", orderUpdated))

		case strings.Contains(sns.TopicArn, "products-event-product-created"):
			var productEvent contract.ProductEvent
			if err := json.Unmarshal([]byte(sns.Message), &productEvent); err != nil {
				return err
			}

			product := productEvent.ToDomain()

			productCreated, err := h.productService.Create(ctx, product)
			if err != nil {
				return err
			}

			zap.L().Info("new product created successfully", zap.Any("product", productCreated))

		case strings.Contains(sns.TopicArn, "products-event-product-updated"):
			var productEvent contract.ProductEvent
			if err := json.Unmarshal([]byte(sns.Message), &productEvent); err != nil {
				return err
			}

			productFound, err := h.productService.FindByPublicID(ctx, productEvent.PublicID)
			if err != nil {
				return err
			}

			product := productEvent.ToDomain()

			productUpdated, err := h.productService.UpdateByID(ctx, productFound.ID, product)
			if err != nil {
				return err
			}

			zap.L().Info("product updated successfully", zap.Any("product", productUpdated))

		case strings.Contains(sns.TopicArn, "products-event-product-deleted"):
			var productEvent contract.ProductEvent
			if err := json.Unmarshal([]byte(sns.Message), &productEvent); err != nil {
				return err
			}

			productFound, err := h.productService.FindByPublicID(ctx, productEvent.PublicID)
			if err != nil {
				return err
			}

			productDeleted, err := h.productService.DeleteByID(ctx, productFound.ID)
			if err != nil {
				return err
			}

			zap.L().Info("product deleted successfully", zap.Any("product", productDeleted))

		}
	}

	return nil
}
