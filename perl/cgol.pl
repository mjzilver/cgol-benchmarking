#!/usr/bin/perl

use strict;
use warnings;
use Getopt::Long;

my $input_file;
my $amount;
my $silent = '';
my $ouput_file = "./output.txt";

GetOptions ("amount=i" => \$amount, 
            "file=s"   => \$input_file, 
            "silent"  => \$silent)
or die("Error in command line arguments\n");

if (!defined $input_file || !defined $amount) {
    print "Missing --file or --amount\n";
    die;
}

my @board;
my $size = -1;

# parse file
sub parse_file {
    open(FH, '<', $input_file) or die "Cannot open $input_file: $!\n";

    my $lines = 0;

    while(<FH>){
        my @chars = split('', $_);

        if ($size == -1) {
            $size = scalar @chars - 1;
        }

        for (my $i = 0; $i < $size; $i++) {
            $board[$lines][$i] = $chars[$i]
        }

        $lines++;
    }

    close(FH);
}

sub print_board {
    my $output = '';

    for (my $y = 0; $y < $size; $y++) {
        for (my $x = 0; $x < $size; $x++) {
            $output .= $board[$y][$x];
        }
        $output .= "\n";
    }

    open(my $fh, '>', $ouput_file) or die "Cannot open $ouput_file: $!\n";  
    print $fh $output;  
    close($fh)
}


my @neighbors = (
    [-1, -1], [-1, 0], [-1, 1],
    [0, -1], [0, 1],
    [1, -1], [1, 0], [1, 1]
);

sub next_state {
    my @next_board;

    for (my $y = 0; $y < $size; $y++) {
        for (my $x = 0; $x < $size; $x++) {
            my $bool = should_cell_live($y, $x);

            if ($bool) {
                $next_board[$y][$x] = '1';
            } else {
                $next_board[$y][$x] = '0';
            }
        }
    }

    return @next_board;
}

sub should_cell_live {
    my ($y, $x) = @_;
    my $count = 0;

    foreach my $offset (@neighbors) {
        my ($dy, $dx) = @$offset;

        # wrapping
        my $ny = ($y + $dy + $size) % $size;
        my $nx = ($x + $dx + $size) % $size;

        if ($board[$ny][$nx]) {
            $count++;
        }
    }


    if ($board[$y][$x]) {
        return 0 if $count < 2 || $count > 3;
    } else {
        return 1 if $count == 3; 
    }

    return $board[$y][$x];  
}

sub main_loop {
    parse_file();

    for (my $y = 0; $y < $amount; $y++) {
        @board = next_state();
    }

    if (!$silent) {
        print_board();
    }
}

main_loop();
