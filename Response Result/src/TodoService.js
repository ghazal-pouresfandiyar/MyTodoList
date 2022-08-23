import http from "./Http-comman";  
class TodoService {   
    getAll() {  
        return http.get("/api/todo"); 
      }        
}  
export default new TodoService(); 