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
        s:Chars,
    }

    let mut dq = VecDeque::new();
    for c in s {
        dq.push_back(c);
    }

    for _ in 0..n {
        let l = dq.len();
        if l == 0 {
            println!("{}", 0);
            return;
        }

        dq.rotate_left(1);
        if l >= 3 {
            if dq[l - 3] == 'f' && dq[l - 2] == 'o' && dq[l - 1] == 'x' {
                dq.truncate(l - 3);
            }
        }
    }
    println!("{}", dq.len());
}
