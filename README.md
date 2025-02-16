# concurrency-problems-golang
This repo contains concurrency problems solved using Go langauge. here is the quick reference for the list of problems solved in this repo -

# [1. Thread-safe Bounded Blocking Queue](https://github.com/maniram-yadav/concurrency-problems-golang/tree/main/blocking-queue)
* A queue with a fixed capacity that supports concurrent enqueue and dequeue operations.
* Features:
  * Blocks enqueue operations when the queue is full.
  * Blocks dequeue operations when the queue is empty.
  * Optimized for high-concurrency use cases.
* Use Cases:
  * Task queues in multi-threaded environments.
  * Producer-consumer problems.
* Implementation:
  * Use channels and synchronization primitives like sync.Mutex and sync.Cond.
# [2. Multi-threaded Reader-Writer Lock](https://github.com/maniram-yadav/concurrency-problems-golang/tree/main/threaded-rwlcok)
* A lock that allows multiple readers or a single writer at any time.
* Features:
  * Prioritize writers or readers based on the use case.
  * Avoid starvation.
* Use Cases:
  * Resource-heavy read operations.
  * Shared configuration management.
* Implementation:
  * Use sync.RWMutex and custom logic to add prioritization.
# [3. Thread Pool with Dynamic Resizing](https://github.com/maniram-yadav/concurrency-problems-golang/tree/main/thread-pool)
* A thread pool that can resize dynamically based on the workload.
* Features:
  * Create worker goroutines on demand.
  * Recycle idle workers.
* Use Cases:
  * Background task execution.
  * High-throughput systems.
* Implementation:
  * Use a combination of worker queues and channels with a monitor goroutine for resizing.
# [4. Dining Philosophers Problem](https://github.com/maniram-yadav/concurrency-problems-golang/blob/main/dining-philosophers)
* Classical synchronization problem with philosophers competing for forks (resources).
* Features:
  * Deadlock-free design.
  * Avoid starvation.
* Use Cases:
  * Resource allocation in distributed systems.
* Implementation:
  * Use semaphores or channels for fork availability.
# [5. Rate Limiter for API Requests](https://github.com/maniram-yadav/concurrency-problems-golang/tree/main/rate-limiter)
* Limits the number of requests over a time window.
* Features:
  * Token bucket or leaky bucket algorithms.
  * Configurable limits per user or service.
* Use Cases:
  * Preventing abuse of APIs.
* Implementation:
  * Use time-based counters and channels for efficient limiting.
# [6. Concurrent File Processor](https://github.com/maniram-yadav/concurrency-problems-golang/blob/main/file-processor)
* Processes multiple files concurrently while maintaining order or dependencies.
* Features:
  * Support for different file formats and operations.
  * Handle errors gracefully.
* Use Cases:
  * Log processing.
  * Data pipeline systems.
* Implementation:
  * Use worker pools and channels for file tasks.
# [7. Job Scheduler with Dependency Management](https://github.com/maniram-yadav/concurrency-problems-golang/tree/main/jobscheduler)
* A scheduler that respects dependencies between jobs.
* Features:
  * DAG-based dependency resolution.
  * Concurrency for independent jobs.
* Use Cases:
  * Build systems.
  * Workflow automation.
* Implementation:
  * Use a DAG structure and worker pools.
