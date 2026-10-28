#!/usr/bin/env python3.14
import os
import shutil
import subprocess

REPO_URL = "https://neutron.cdm.depaul.edu/csc-376/assignments/csc-376-auto-grader.git"

def main():
    # Clone the auto-grader repository into the test folder
    clone_dir = os.path.join('test', 'auto-grader')
    if os.path.exists(clone_dir):
        print(f"Updating existing repository at {clone_dir}...")
        subprocess.run(["git", "pull"], cwd=clone_dir, check=True)
    else:
        print(f"Cloning {REPO_URL} into {clone_dir}...")
        subprocess.run(["git", "clone", REPO_URL, clone_dir], check=True)

    # Copy test/traces folder recursively from the cloned repo to the current repo
    src_traces_dir = os.path.join(clone_dir, 'test', 'traces')
    dst_traces_dir = os.path.join('test', 'traces')
    if os.path.exists(src_traces_dir):
        if os.path.exists(dst_traces_dir):
            shutil.rmtree(dst_traces_dir)
            print(f"Overwriting existing {dst_traces_dir}")
        shutil.copytree(src_traces_dir, dst_traces_dir)
        print(f"Copied {src_traces_dir} to {dst_traces_dir}")
    else:
        print(f"Error: {src_traces_dir} not found.")

    # Copy test/test_driver_pa5.py from the cloned repo to the root of the current repo
    src_test_file = os.path.join(clone_dir, 'test', 'test_driver_pa5.py')
    dst_test_file = os.path.join('test', 'test_driver.py')
    if os.path.exists(src_test_file):
        if os.path.exists(dst_test_file):
            os.remove(dst_test_file)
            print(f"Overwriting existing {dst_test_file}")
        shutil.copy2(src_test_file, dst_test_file)
        print(f"Copied {src_test_file} to {dst_test_file}")
    else:
        print(f"Error: {src_test_file} not found.")

if __name__ == "__main__":
    main()