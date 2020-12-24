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
        t:usize,
        s:[Chars;t],
    }

    for v in s {
        if let Some(p) = v.iter().position(|&c| c != 'a') {
            if p == 0 {
                println!("{}", 0);
                continue;
            }

            let mut v_r = v.clone();
            v_r[1] = v[p];
            if v_r.into_iter().collect::<String>() > "atcoder".to_string() {
                println!("{}", p - 1);
            } else {
                println!("{}", p);
            }
        } else {
            println!("{}", -1);
        }
    }
}
