class Set:
    def __init__(self, name="Anon"):
        self.items = {}
        self.name = name

    def __repr__(self):
        return f"{self.name}: {self.items}"

    def __iter__(self):
        for item in self.items:
            yield item

    def __len__(self):
        return len(self.items)

    def add(self, item):
        if item in self.items:
            return False
        self.items[item] = item
        return True
