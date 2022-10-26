from functools import wraps


def log_middleware():
    def _log_middleware(f):
        @wraps(f)
        def __log_middleware(*args, **kwargs):
            print('before')
            result = f(*args, **kwargs)
            print('result: %s' % result)
            return result
        return __log_middleware
    return _log_middleware
