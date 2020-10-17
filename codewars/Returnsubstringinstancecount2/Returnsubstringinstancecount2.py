import re
def search_substr(full_text, search_text, allow_overlap=True):
    if not (full_text and search_text): return 0
    return len(re.findall(f"({allow_overlap and '?=' or ''}{search_text})", full_text))