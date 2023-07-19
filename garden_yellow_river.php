<?php 
//Initializing Variables 
$title = "Creative Refresh"; 
$name = "Some Name";
$version = "1.0";

//Function to Write a Header 
function writeHeader() { 
	global $title; 
	echo "<h1>".$title."</h1>";
}

//Function to Retrieve the Name 
function getName() { 
	global $name; 
	return $name; 
}

//Function to Retrieve the Version 
function getVersion() { 
	global $version; 
	return $version; 
}

//Function to Capture the Data
function captureData() { 
	$data = array();
	$data['name'] = getName(); 
	$data['version'] = getVersion(); 
	return $data; 
}

//Function to Generate the HTML
function generateHTML() { 
	$data = captureData(); 
	$html = "<p>Name: ".$data['name']."</p>"; 
	$html .= "<p>Version: ".$data['version']."</p>"; 
	echo $html; 
}

//Function to Process a Request 
function processRequest() { 
	writeHeader(); 
	generateHTML();
}

//Main Program 
processRequest(); 

//Function to Generate the CSS 
function generateCSS() { 
	echo "h1 { font-size: 30px; font-family: Arial; }"; 
}

//Function to Generate the JavaScript 
function generateJS() { 
	echo "document.write('<h1>Creative Refresh</h1>');"; 
}

//Function to Generate the Markup 
function generateMarkup() { 
	echo "<br/><div><h2>Hello World!</h2></div>"; 
}

//Function to Generate the Resources 
function generateResources() { 
	generateCSS(); 
	generateJS(); 
	generateMarkup(); 
}

//Function to Draw the Resources 
function drawResources() { 
	$resources = generateResources(); 
	echo "<style>".$resources[0]."<style>";
	echo "<script>".$resources[1]."</script>"; 
	echo $resources[2]; 
}

//Main Program 
drawResources(); 
?>