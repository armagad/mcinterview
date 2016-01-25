//Create your form validation here.//

REQUIRED_FIELDS = {
	first_name: function(e) { if(e.target.value == '') return "First Name is required."},
	last_name: function(e) { if(e.target.value == '') return "Last Name may not be blank."},
	email: function(e) { if(e.target.value == '') return "Email is required."},
	message:  function(e) { if(e.target.value == '') return "Short Bio is required."; if(e.target.value.length > 1000) return "Character Limit: 1000 characters."}
};

document.getElementById("partI").onsubmit = function(){
	var keys = Object.keys(REQUIRED_FIELDS);
	var isReady = true;

	for(var i = 0; i < keys.length; i++){
		var el  = document.getElementById(keys[i]);
		var label = el.parentNode;

		label.title = '';
		label.style.color = 'black';
		label.style.textDecoration = 'none';

		var err = REQUIRED_FIELDS[keys[i]]({target: el});
		if(err){
			isReady = false;
			label.title = err;
			label.style.color = 'yellow';
			label.style.textDecoration = 'underline';
		}
	}

	return isReady;
}
