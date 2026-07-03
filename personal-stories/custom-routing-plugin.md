***Business Context***

    We had APIs deployed across multiple regions and backend services. We needed to route requests dynamically based on request attributes such as headers, query parameters, and path parameters.

***Pain Point***

The existing approach relied on a separate service that stored routing configurations and determined where requests should be sent.

This introduced:
- Additional network hops
- Operational overhead
- A central dependency for routing decisions
- Difficulty scaling routing logic across teams and environments

Kong's native routing capabilities were not flexible enough for our use case.

***My Responsibility***

I designed and implemented a Kong plugin that enabled teams to define routing rules declaratively instead of building custom routing logic into services.

***Solution***

The plugin allowed teams to configure routing rules using YAML.

Rules could match on:
- Headers
- Query parameters
- Path segments

Based on the matched rule, the plugin could:

- Select a backend service
- Rewrite request paths
- Apply fallback rules
- Prioritize rule execution

This moved routing decisions directly into the gateway layer.

***Technical Decisions***

This is where interviewers often dig deeper.

I would prepare answers for:

Why a plugin instead of another service?

Avoid extra network calls.
Keep routing decisions close to request processing.
Reduce operational complexity.

Why configuration-driven instead of code-driven?

Teams could change routing behavior without code deployments.
Reusable across multiple products and environments.

How did you prevent bad configurations?

Configuration validation before activation.
Restricted allowed template substitutions.
Controlled regex usage.
Prevented malformed upstream URLs.

How did you handle rule evaluation?

Priority ordering.
First-match-wins strategy.
Default fallback route.

***Challenges***

This section can be much stronger.

Instead of:

Identifying the level of flexibility to give.

Say:

The biggest challenge was balancing flexibility with safety. Teams wanted powerful matching and rewrite capabilities, but excessive flexibility could create security risks, routing loops, or malformed requests. We had to carefully define what was configurable and validate configurations before they were applied.

That sounds like senior engineering thinking.

Impact

I would phrase it as:

The plugin eliminated the need for a dedicated routing-configuration service, reduced operational complexity, and enabled application teams to own routing behavior through configuration rather than requiring gateway changes for every new routing rule.

Even if you don't have hard metrics, that's a solid impact statement.

What I Learned

Your answer is actually good.

I'd say:

It taught me how to build framework-level solutions. Instead of solving a single routing problem, I had to design a reusable capability that different teams could use safely without involving the platform team for every change.


Why couldn't Kong routes solve this?
Why a plugin instead of a microservice?
How were rules stored and loaded?
How did you validate configurations?
What security risks did you consider?
How did you handle performance impact on every request?