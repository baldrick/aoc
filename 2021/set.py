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

    def __getitem__(self, item):
        if item in self.items:
            return self.items[item]
        return None

    def add(self, item):
        if item in self.items:
            return False
        self.items[item] = item
        return True

    def remove(self, item):
        del self.items[item]