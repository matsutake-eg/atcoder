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
        s:Chars,
    }

    if s.len() == 1 {
        if s[0] == '8' {
            println!("Yes");
        } else {
            println!("No");
        }
        return;
    } else if s.len() == 2 {
        if s.iter().collect::<String>().parse::<usize>().unwrap() % 8 == 0 {
            println!("Yes");
        } else if s.iter().rev().collect::<String>().parse::<usize>().unwrap() % 8 == 0 {
            println!("Yes");
        } else {
            println!("No");
        }
        return;
    }

    let mut hm = HashMap::new();
    for c in s {
        *hm.entry(c).or_insert(0) += 1;
    }

    let mut i = 104;
    while i < 1000 {
        let mut tm = HashMap::new();
        for c in format!("{:03}", i).chars() {
            *tm.entry(c).or_insert(0) += 1;
        }
        let mut is_ok = true;
        for (k, v) in tm {
            if hm.get(&k) == None || hm.get(&k) < Some(&v) {
                is_ok = false;
                break;
            }
        }
        if is_ok {
            println!("Yes");
            return;
        }
        i += 8;
    }
    println!("No");
}
