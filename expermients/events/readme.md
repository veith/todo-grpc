Ein Modul kennt seine Handler und Events mit ihren definierten Typen.

Über die Konfiguration registirert ein Modul seine Handler zu den entsprechenden Topics und 
wird die Events auf die entsprechenden Topics werfen. Topics auf die ein Modul wirft beginnen immer mit dem Modulnamen selbst.
Ein Modul **Task** kann nur auf ein Topic werfen das mit **task.** beginnt. Folgende Namen für Topics sind gültig:
- **task.created**
- **task.vip.created**
- **task.vip.completed**

 


![flow](flow.png)
