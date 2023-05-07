package ejercicios


//1. Implementar el TAD Conjunto sobre un Map de Go.
// Tal que las operaciones Add, Remove y Search sean O(1). 
//Sabiendo que los Map de Go usan hashing para almacenar sus claves.
// La idea por lo tanto es aprovechar el map para almacenar los elementos del conjunto como las claves de un map
type Set [T comparable] struct{
  elementos Map[T]{}
 }
