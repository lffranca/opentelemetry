from flask import abort, jsonify
from flask_restful import Resource
from opentelemetry import trace, metrics

from src.flaskr.models import Product

tracer = trace.get_tracer(__name__)
meter = metrics.get_meter(__name__)

# Now create a counter instrument to make measurements with
product_get_all_counter = meter.create_counter(
    "resource.product.get_all",
    description="The number of request",
)


class ProductResource(Resource):
    def get(self):
        with tracer.start_as_current_span("products_get_all") as span:
            products = Product.query.all() or abort(204)
            span.set_attribute("products_get_all.value", products)
            # This adds 1 to the counter for the given roll value
            product_get_all_counter.add(1)
            return jsonify(
                {"products": [product.to_dict() for product in products]}
            )


class ProductItemResource(Resource):
    def get(self, product_id):
        with tracer.start_as_current_span("products_get") as span:
            product = Product.query.filter_by(id=product_id).first() or abort(404)
            span.set_attribute("products_get.value", product)
            return jsonify(product.to_dict())
