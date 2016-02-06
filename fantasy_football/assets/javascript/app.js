
(function( pjr, $, undefined ) {
    // Private Property
    // var private = "property;

    // Public Property
    // pjr.public = "property";
    pjr.error_mapping = {
      'reg_err_0' : ['er_ps'],
      'reg_err_1' : ['er_ps'],
      'reg_err_2' : ['er_un', 'er_em'],
      'reg_err_3' : ['er_em'],
      'log_err_0' : ['er_un'],
    };
    pjr.error_messages = {
      'reg_err_0' : 'password minimum length is 7',
      'reg_err_1' : "passwords don't match",
      'reg_err_2' : 'email/username already taken',
      'reg_err_3' : 'invalid email',
      'log_err_0' : 'invalid login',
    };
    //Public Method
    pjr.addEvent = function(object, type, callback){
      if (object === null || typeof(object) === 'undefined') return;
      if (object.addEventListener) {
        object.addEventListener(type, callback, false);
      } else if (object.attachEvent) {
        object.attachEvent("on" + type, callback);
      } else {
        object["on"+type] = callback;
      }
    };

    pjr.isEmpty = function (obj) {
      if (obj == null )      return true;
      if (obj.length > 0)    return false;
      if (obj.length === 0)  return true;
      for (var key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) return false;
      }
      return true;
    };

    pjr.parseForm = function(form){
      var elements = form.elements;
      var data = {};
      var errors = {};
      for (var i = 0, element; element = elements[i++];) {
        if (element.type === "button")
          ;
        else if ((element.value === "" || element.value === null) && element.required === true)
          errors[i] = element.id;
        else{
          var key = element.name;
          data[key] = element.value;
        }
      }
      if (!isEmpty(errors)) {
        return {success: false, data: errors };
      }
      return {success: true, data: data };
    };

    pjr.clearFormFields = function(form){
      var elements = form.elements;
      for (var i = 0, element; element = elements[i++];) {
        if (element.type === "button")
          ;
        else{
          element.value = "";
        }
      }
    };

    pjr.sendForm = function(form, url, callback){
      $.ajax({
        type: "POST",
        url: url,
        data: $(form).serialize(), // serializes the form's elements.
        success: function(data){
          callback(data)
        }
      });
    };

    pjr.applyErrorMessages = function (message, map){
      if (map.hasOwnProperty(message)){
        map[message].forEach(function(value){
          var err_msg = ['  *(', pjr.error_messages[message], ')'];
          document.getElementById(value).innerHTML = err_msg.join('');
        });
      }
    };

    // Private Method
    // function addItem( item ) {
    //     if ( item !== undefined ) {
    //         console.log( "Adding " + $.trim(item) );
    //     }
    // }
    
}( window.pjr = window.pjr || {}, jQuery ));
