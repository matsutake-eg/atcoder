#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        r1:i64,
        c1:i64,
        r2:i64,
        c2:i64,
    }

    let r = r2 - r1;
    let c = c2 - c1;
    if r == 0 && c == 0 {
        println!(0);
        return;
    }
    if r == c || r == -c {
        println!(1);
        return;
    }
    if r.abs() + c.abs() <= 3 {
        println!(1);
        return;
    }
    if r.abs() + c.abs() <= 6 {
        println!(2);
        return;
    }
    if (r + c).abs() <= 3 || (r - c).abs() <= 3 {
        println!(2);
        return;
    }
    if (r ^ c ^ 1) & 1 == 1 {
        println!(2);
        return;
    }
    println!(3);
}
