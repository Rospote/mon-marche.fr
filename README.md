**Solution proposée**

On accepte le payload en tant que rawData sur le endpoint /ticket.  
On convertit l'entrée en tant que string. Pour la suite, on va le parse ligne par ligne.  
On vérifie que l'input est valide.  
Si l'input est valide, on convertit la donnée en objets TicketHeader et Products.  
On fait ensuite un insert en DB pour le ticketHeader et la liste des produits.  
Si l'input n'est pas valide, on stocke le ticket tel quel en BDD  

**Choix**

Le traitement de l'input en tant que string pour plus de flexibilité et de contrôle sur l'attendu via des regex  
La gestion des erreurs d'input se fait en entrée de l'algorithme pour anticiper au plus vite les problèmes  
Typage fort des données pour protéger la BDD des insertions de données non voulues  
Les tickets non conformes sont gérés en BDD pour faciliter le stockage et le traitement ultérieur  
Les requêtes des produits sont parallélisées pour optimiser rapidement le traitement de ceux ci  
Découpage en fichiers en fonction des concerns de chacun.  

Modele de données :
Les tables en BDD sont:  
-tblTicketHeaders(order_id (PK),vat, total)  
-tblProducts(id(PK), fk_orderid, productname, price)  
-tblErrorTickets(id(PK),ticket)  

**Pistes d'amélioration**

La vérification sur l'input n'est pas totalement fault tolerant et doit être amélioré  
La solution ne prend actuellement pas en charge le cas où le produit est déjà en BDD  
Certains tickets peuvent actuellement être perdus (par ex: erreur a l'insertion en BDD). Il faut envisager un mécanisme de récupération des tickets supplémentaire en dehors de la BDD  
Les requêtes en BDD peuvent être optimisées (bulk pour l'insertion des produits)  
On pourrait implementer une solution de multithreading plus sophistiquée (ex: worker threads, à dimensionner selon les capacités de la machine)  
--> Traitement en chunks envisageable en fonction de la volumétrie. Dans ce cas revoir les appels vers la BDD  
La charge en entrée pourrait être gérée par un load balancer + scaling horizontal du web service  
