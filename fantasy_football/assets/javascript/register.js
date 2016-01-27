// <form method='POST' action = '/register' id ='reg_form' >
//     <input type="text" name="user[name]" id="register_user_name" required/>
//     <input type="text" name="user[email]" id="register_email" required/>
//     <input type="password" name="user[password]" id="register_password" required/>
//     <input type="password" name="user[confirm_password]" id="register_confirm_password" required/>
//     <button type="submit" id="register_button" >register</button>
// </form>
// var register_button = document.getElementById("register_button");
// var reg_form        = document.getElementById("reg_form");

// pjr.addEvent(register_button, "click", validate_register_form);

// var validate_register_form = function(){
//   var data = pjr.parseForm(reg_form)
//   if (data.success){
//     loadingScreen("Submitting Registration...");
//     pjr.sendForm(reg_form, "/register", register_form_onsubmit);
//   }
// }

// var register_form_onsubmit = function(){
// }

// var loadingScreen = function(){
  
// }


// clearFormFields