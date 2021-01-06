#!/usr/bin/python3
from libs import stuff
from os import getcwd, getenv, chdir
from os import system as syscall
syscall("sudo apt install python3-tk") 
syscall("sudo apt install python3-pil python3-pil.imagetk")
from tkinter import *
from tkinter.ttk import *
import time
from PIL import ImageTk,Image

# creating tkinter window
root = Tk()


make_link = "sudo ln -s %s/.pck3r/core_pck3r.py /bin/pck3r" % getenv("HOME")

# Progress bar widget
progress = Progressbar(root, orient = HORIZONTAL,
length = 100, mode = "determinate")

# Function responsible for the updation
# of the progress bar value
def bar():


    progress["value"] = 20
    root.update_idletasks()
    time.sleep(1)
    print("Unlinking pck3r (if was installed) ")
    syscall("sudo unlink /bin/pck3r ")
    syscall("sudo rm -rf /bin/pck3r*")
    progress["value"] = 40
    root.update_idletasks()
    time.sleep(1)





    syscall("mkdir ~/.pck3r") if (syscall("echo %s ; ls ~/.pck3r" % stuff.CYN)) !=0 else print("%sCopy all pck3r directory %s" %(stuff.CYN, stuff.NRM))
    syscall("cp -rf . ~/.pck3r")
    progress["value"] = 50
    root.update_idletasks()
    time.sleep(1)

    syscall(make_link)
    progress["value"] = 60
    root.update_idletasks()
    time.sleep(1)

    print("%sCheck link " % stuff.YEL)
    print("Link created ") if (syscall("ls -l  /bin/pck3r")) == 0  else print("%s%sNo link (/bin/pck3r)%s " % (stuff.sysOk(), stuff.RED, stuff.NRM))
    progress["value"] = 80
    root.update_idletasks()
    time.sleep(1)
    progress["value"] = 100

    print("%s%sPck3r installed successfuly %s" % (stuff.sysOk(), stuff.GRN, stuff.NRM))
    root.quit()


# user theme and window configuration
root.title("Pck3r Installer")
icon = Image.open("%s/.pck3r/icon/pck3r-logo.png" % getenv("HOME"))
photo = ImageTk.PhotoImage(icon)
root.wm_iconphoto(False, photo)


root.geometry("300x80")
root.configure(background="black")

root.resizable(False, False)
root.style = Style()
root.style.theme_use("alt")


# packing to main window (panel)

progress.pack(fill="x", pady = 10)
Button(root, text="install pck3r (system wide) ", command = bar).pack(fill="x", pady = 10)

# infinite loop
mainloop()
