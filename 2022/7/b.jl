FILE=1
DIR=2

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
    end
    du(root)
    report()
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
        return root
    end
    if newDir == ".."
        return tree.parent
    end
    if ! haskey(tree.dirs, newDir)
        cwd = tree.cwd.name
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
    end
    tree.size = s
end

TOTAL=70000000
NEED=30000000

function report()
    free = TOTAL - root.size
    minDeleteSize = NEED - free
    println("$TOTAL - $(root.size) = $free, min to delete=$NEED - $free = $minDeleteSize")
    min = findMinOver(minDeleteSize, root)
    println("$min")
end

function findMinOver(min, tree)
    if tree.size < min
        # Ignore directories that won't free up enough space
        return 999999999999999
    end
    minimumOver = tree.size
    for (dir, entry) in tree.dirs
        localMin = findMinOver(min, entry)
        if localMin < minimumOver
            minimumOver = localMin
        end
    end
    return minimumOver
end

main()
