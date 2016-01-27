
(function( pjr, $, undefined ) {
    // Private Property
    // var private = "property;

    // Public Property
    // pjr.public = "property";

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

    // Private Method
    // function addItem( item ) {
    //     if ( item !== undefined ) {
    //         console.log( "Adding " + $.trim(item) );
    //     }
    // }
    
}( window.pjr = window.pjr || {}, jQuery ));
