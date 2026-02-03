# CWE-502: Insecure Deserialization - Python

## LLM Guidance

Python's `pickle` module executes arbitrary code during deserialization, enabling remote code execution when unpickling untrusted data. Attackers craft malicious pickle payloads that invoke `__reduce__` or `__setstate__` methods to run system commands. Primary fix: Replace pickle with JSON (`json.loads()`), MessagePack, or Protocol Buffers for all untrusted data.

## Key Principles

- Never use `pickle.loads()`, `pickle.load()`, `marshal.loads()`, `shelve`, or `pd.read_pickle()` with untrusted data
- Replace pickle with JSON for object serialization (requires manual object reconstruction)
- Use `yaml.safe_load()` instead of `yaml.load()` or `yaml.unsafe_load()`
- For pandas DataFrames, use CSV, Parquet, or Feather formats instead of pickle
- If pickle is absolutely required, implement a restricted unpickler with class allowlisting

## Remediation Steps

- Identify all deserialization calls (`pickle.loads()`, `pd.read_pickle()`, `yaml.load()`, etc.)
- Replace with safe alternatives - `json.loads()` for objects, `pd.read_parquet()` for DataFrames
- Update file extensions and storage mechanisms (`.pkl` → `.json` or `.parquet`)
- Manually reconstruct objects from deserialized dictionaries with validation
- For Django sessions, set `SESSION_SERIALIZER = 'django.contrib.sessions.serializers.JSONSerializer'`
- Test that legitimate data flows work correctly with new serialization format

## Safe Pattern

```python
# Replace pickle with JSON
import json

# Serialize
user_dict = {'name': user.name, 'email': user.email}
json_data = json.dumps(user_dict)

# Deserialize safely
data = json.loads(json_data)  # Creates dict, not arbitrary objects
user = User(name=data['name'], email=data['email'])  # Explicit reconstruction
```
