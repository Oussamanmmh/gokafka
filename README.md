# GoKafka

## 🔔 NOTICE

### 🔹 Consumer
Reads messages from a topic.

### 🔹 Consumer Group
A set of consumers that work together to consume a topic without duplicating work.

---

## 🧪 Two Main Scenarios

### ✅ Scenario 1: Multiple Consumers in the Same Consumer Group
- Each message is **delivered to only one** consumer in the group.
- Kafka will **load balance** partitions across them.
- Used for **parallel processing**.

**Example:**  
You have 3 consumers in group `order-processors`. They each process different orders from the topic.

---

### ✅ Scenario 2: Multiple Consumers in Different Consumer Groups
- Each group is **independent**.
- All groups **get all messages** — like a **broadcast**.
- Used when **different systems need to consume the same data**.

**Example:**  
Group `analytics-service` reads all orders for stats.  
Group `order-service` reads all orders for fulfillment.
