#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        k:usize,
        s:Chars,
    }

    let mut s = s;
    for _ in 0..k {
        s = (0..n)
            .map(|i| win(s[i * 2 % n], s[(i * 2 + 1) % n]))
            .collect();
    }
    println!("{}", s[0]);
}

fn win(a: char, b: char) -> char {
    match (a, b) {
        (a, b) if a == 'R' && b == 'P' => b,
        (a, b) if a == 'P' && b == 'S' => b,
        (a, b) if a == 'S' && b == 'R' => b,
        _ => a,
    }
}
