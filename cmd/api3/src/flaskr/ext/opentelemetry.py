# from opentelemetry.instrumentation.flask import FlaskInstrumentor
#
#
# def init_app(app):
#     FlaskInstrumentor().instrument_app(app)
import logging
from opentelemetry import trace, metrics
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import OTLPLogExporter
from opentelemetry.sdk._logs import (
    LogEmitterProvider,
    LoggingHandler,
    set_log_emitter_provider,
)
from opentelemetry.sdk._logs.export import BatchLogProcessor


def init_app(app):
    # Service name is required for most backends
    resource = Resource(attributes={
        SERVICE_NAME: "app_03"
    })

    # TRACER
    tracer_provider = TracerProvider(resource=resource)
    processor = BatchSpanProcessor(OTLPSpanExporter(
        endpoint="localhost:4317",
        insecure=True,
    ))
    tracer_provider.add_span_processor(processor)
    trace.set_tracer_provider(tracer_provider)

    # METER
    reader = PeriodicExportingMetricReader(
        OTLPMetricExporter(
            endpoint="localhost:4317",
            insecure=True,
        )
    )
    meter_provider = MeterProvider(resource=resource, metric_readers=[reader])
    metrics.set_meter_provider(meter_provider)

    # LOG
    log_emitter_provider = LogEmitterProvider(
        resource=resource,
    )
    set_log_emitter_provider(log_emitter_provider)

    exporter = OTLPLogExporter(
        endpoint="localhost:4317",
        insecure=True,
    )
    log_emitter_provider.add_log_processor(BatchLogProcessor(exporter))
    handler = LoggingHandler(
        level=logging.NOTSET, log_emitter_provider=log_emitter_provider
    )

    # Attach OTLP handler to root logger
    logging.getLogger().addHandler(handler)
