#!/usr/bin/perl

#Creative Refresh

#Creates a new file that is populated with 1000 random alphanumeric characters

use strict;
use warnings;

my @chars = ('a'..'z', 'A'..'Z', 0..9);

open (my $outputFile, '>', 'RandomCharacters.txt') or die "Could not open file 'RandomCharacters.txt': $!";

for (my $i = 0; $i < 1000; $i++) {
  	my $character = $chars[rand @chars];
  	print $outputFile "$character";
}

close $outputFile;

#Opening the file containing 1000 random alphanumeric characters

open(my $inputFile, '<', 'RandomCharacters.txt') or die "Could not open file 'RandomCharacters.txt': $!";

#Creating a new file containing 2000 random alphanumeric characters

open(my $finalOutputFile, '>', 'NewRandomCharacters.txt') or die "Could not open file 'NewRandomCharacters.txt': $!";

while (my $line = <$inputFile>) {
	for (my $i = 0; $i < 1000; $i++) {
  		my $character = $chars[rand @chars];
  		print $finalOutputFile "$character";
	}
	print $finalOutputFile $line;
}

close $inputFile;
close $finalOutputFile;

#Removing the file with 1000 characters

unlink 'RandomCharacters.txt' or die "Could not remove file 'RandomCharacters.txt': $!";