//Create your form validation here.//

REQUIRED_FIELDS = {
	first_name: {label: "First Name", timeout: null, tests: []},
	last_name:  {label: "Last Name",  timeout: null, tests: []},
	email:      {label: "Email",      timeout: null, tests: [el => asyncEmailTeset(el)]},
	message:    {label: "Short Bio",  timeout: null, tests: [el => (el.value.length > 1000)? "Character Limit: 1000 characters.": ""]}
};

function handleChange(e){
	clearTimeout(REQUIRED_FIELDS[e.target.id].timeout);
	REQUIRED_FIELDS[e.target.id].timeout = setTimeout( () => handleBlur(e), 750);
}

function handleBlur(e){

	var isClear = true;
	REQUIRED_FIELDS[e.target.id].tests
		.map( fn => fn(e.target) )
		.forEach( function(msg) { warnRequired(e.target.id, msg) ; isClear = false } ) ;

	if(e.target.value == ''){
		warnRequired(e.target.id, REQUIRED_FIELDS[e.target.id].label + " is required.");
		isClear = false;
	}

	if(isClear){
		var label = e.target.parentNode;

		label.title = '';
		label.style.color = 'black';
		label.style.textDecoration = 'none';
		
		isAllClear();
	}
}

function asyncEmailTeset(el){

	var r = new XMLHttpRequest();
	r.open('GET', 'http://codetest.socrative.com/v1/api/email-check/' + el.value + '/');
	r.onreadystatechange = function(){
		if(r.readyState == 4){
			if(r.status == 409){warnRequired(el.id, "Email already exists.")}
			if(r.status == 400){warnRequired(el.id, "Invalid Email.")}
		}
	};
	r.send();
	
	return '';
}

function warnRequired(id, msg){
	if(msg == '') return;

	var el = document.getElementById(id);
	var label = el.parentNode;

	label.title = msg;
	label.style.color = 'yellow';
	label.style.textDecoration = 'underline';

	document.getElementById("submit").disabled = true;
}

function isAllClear(){
	var keys = Object.keys(REQUIRED_FIELDS);

	for(var i = 0; i < keys.length; i++){

		var label = document.getElementById(keys[i]).parentNode;

		if(label.style.color == 'yellow'){ // hacky, I know
			return false;
		}
	}
	
	document.getElementById("submit").disabled = false;
	return true;
}

// Initialize

var keys = Object.keys(REQUIRED_FIELDS);
for(var i = 0; i < keys.length; i++){
	var el  = document.getElementById(keys[i]);
	
	el.onkeyup = handleChange;
	el.onblur  = handleBlur;
}

// Still need backend validation, rely on it
//
// document.getElementById("partI").onsubmit = function(){
// 	var keys = Object.keys(REQUIRED_FIELDS);
// 	var isReady = true;
//
// 	for(var i = 0; i < keys.length; i++){
// 		var el  = document.getElementById(keys[i]);
// 		var label = el.parentNode;
//
// 		label.title = '';
// 		label.style.color = 'black';
// 		label.style.textDecoration = 'none';
//
// 		var err = REQUIRED_FIELDS[keys[i]]({target: el});
// 		if(err){
// 			isReady = false;
// 		}
// 	}
//
// 	return isReady;
// }
