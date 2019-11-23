 const api = (api,user) =>{
    //  console.log(user.us)
    var data = JSON.stringify(user);
      
      var xhr = new XMLHttpRequest();
    //   xhr.withCredentials = true;
      
    //   xhr.addEventListener("readystatechange", function () {
    //     if (this.readyState === 4) {
    //       console.log(this.responseText);
    //     }
    //   });
      
      xhr.open("POST", "/signup");
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.setRequestHeader("cache-control", "no-cache");
      
      xhr.send(data);

 }
 export default api
