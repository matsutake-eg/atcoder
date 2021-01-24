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
        n:usize,
        ab: [(usize,usize);n],
    }

    let mut abc = vec![(0usize, 0usize, 0usize); n];
    let mut ao = 0;
    for i in 0..n {
        let (a, b) = ab[i];
        abc[i] = (a, b, a * 2 + b);
        ao += a;
    }
    abc.sort_by(|a, b| (&b.2).cmp(&a.2));

    let mut tk = 0;
    for i in 0..n {
        let (a, b, _) = abc[i];
        tk += a + b;
        ao -= a;
        if tk > ao {
            println!("{}", i + 1);
            return;
        }
    }
}
