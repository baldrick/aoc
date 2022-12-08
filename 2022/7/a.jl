FILE=1
DIR=2

sum = 0

struct FsEntry
    name::String
    size::Int
    type::Int
end

mutable struct Tree
    cwd::FsEntry
    parent::Union{Tree, Nothing}
    dirs::Dict{String, Tree}
    files::Array{FsEntry}
    size::Int
end

root = Tree(FsEntry("/", 0, DIR), nothing, Dict(), [], 0)

function main()
    f = open(ARGS[1], "r")
    tree = root
    while ! eof(f)
        #=
        $ cd / or .. or directory
        $ ls
        <size> file or dir directory
        =#
        line = readline(f)
        tree = process(tree, line)
        cwd = tree.cwd.name
        println("at $cwd")
    end
    du(root)
    report(root)
    println("$sum")
end

function process(tree, line)
    if startswith(line, "\$")
        tree = processCmd(tree, line[3:end])
    else
        tree = processFsEntry(tree, line)
    end
    return tree
end

function processCmd(tree, cmd)
    if cmd == "ls"
        return tree
    end
    newDir = cmd[4:end]
    if newDir == "/"
        println("cd to root")
        return root
    end
    if newDir == ".."
        return tree.parent
    end
    if ! haskey(tree.dirs, newDir)
        cwd = tree.cwd.name
        println("Add $newDir to $cwd")
    tree.dirs[newDir] = Tree(FsEntry(newDir, 0, DIR), tree, Dict(), [], 0)
    end
    return tree.dirs[newDir]
end

function processFsEntry(tree, line)
    fse = decodeFsEntry(line)
    if fse.type == FILE
        push!(tree.files, fse)
    else # DIR
        if ! haskey(tree.dirs, fse.name)
            cwd = tree.cwd.name
            println("Add $fse to $cwd")
            tree.dirs[fse.name] = Tree(fse, tree, Dict(), [], 0)
        end
    end
    return tree
end

function decodeFsEntry(line)
    s = split(line)
    if s[1] == "dir"
        return FsEntry(s[2], 0, DIR)
    end
    return FsEntry(s[2], parse(Int, s[1]), FILE)
end

function du(tree)
    s = 0
    for (dir, entry) in tree.dirs
        du(entry)
        s += entry.size
    end
    for fse in tree.files
        f = fse.name
        fs = fse.size
        s += fs
        println("$f: $fs")
    end
    tree.size = s
end

function report(tree)
    if tree.size <= 100000
        cwd = tree.cwd.name
        s = tree.size
        global sum += s
        println("$cwd = $s")
    end
    for (dir, entry) in tree.dirs
        report(entry)
    end
end

main()
