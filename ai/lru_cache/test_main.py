from ai.lru_cache.main import LRUCache


def test_cache():
    capacity = 10
    obj = LRUCache(capacity)
    param_1 = obj.get(1)
    obj.put(1, 1)
    obj.put(2, 2)
    param_2 = obj.get(1)
    assert param_1 == -1
    assert param_2 == 1
    