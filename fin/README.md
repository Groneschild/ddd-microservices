# what do?
- make(calculate) and approve budgets
- send invoices

## Connections to other Bounded Contexts (Events)
- ProductSelectie (vanuit ZorgtechSelectie)
    Signaleert de financiering dat er een product is uitgekozen die de patient nodig heeft en begint het proces van Bugdets aanvragen, evalueren and goedkeuren.
- BudgetGoedgekeurt (naar Implementatie)
    Signaleert de Implementatie dat het budget goedgekeurt is en dat het plan uitgevoerd kan worden.
- ImplementatieVoltooit (vanuit Implementatie)
    Signaleert dat de Implementatie voltooit is en dat de Factuur betaald moet zijn. (IDK why invoice is after the product is sent out but whatever)