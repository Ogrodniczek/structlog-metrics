from time import sleep
import structlog
import logging
import orjson


structlog.configure(
    cache_logger_on_first_use=True,
    wrapper_class=structlog.make_filtering_bound_logger(logging.INFO),
    processors=[
        structlog.threadlocal.merge_threadlocal_context,
        structlog.processors.add_log_level,
        structlog.processors.format_exc_info,
        structlog.processors.TimeStamper(fmt="iso", utc=False),
        structlog.processors.JSONRenderer(serializer=orjson.dumps),
    ],
    logger_factory=structlog.BytesLoggerFactory(),
)


logger = structlog.get_logger()
log = logger.bind()

while 1:
    structlog.get_logger("test").warning("hello")
    sleep(0.5)
