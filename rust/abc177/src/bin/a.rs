use proconio;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        d:u32,
        t:u32,
        s:u32,
    }

    println!("{}", if t * s >= d { "Yes" } else { "No" });
}
