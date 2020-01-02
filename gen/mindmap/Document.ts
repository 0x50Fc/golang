import { Id } from "./lib";
import { Element } from "./Element";
import { EventEmitter } from "./Event";
import { CommandQueue, Command } from './Command';

interface ElementSet {
    [id: number]: Element
}

interface DataItem {
    [key: string]: any
    $id: Id,
    $pid?: Id
}

interface Data {
    id: Id
    items: DataItem[]
}

export class Document extends EventEmitter {


    get root(): Element {
        return this._root
    }

    protected _root: Element
    protected _id: Id
    protected _elementSet: ElementSet
    protected _queue: CommandQueue | undefined

    constructor(queue?: CommandQueue) {
        super();
        this._queue = queue;
        this._elementSet = {}
        this._id = 0;
        this._root = new Element(this, 0);
    }

    save(): Data {

        let items: DataItem[] = [];

        function add(e: Element) {

            let item: DataItem = {
                $id: e.id,
                $pid: e.parent !== undefined ? e.parent.id : undefined
            }

            if (typeof e.data == 'object') {
                for (let key in e.data) {
                    if (key.length > 0 && key.substr(0, 1) == '$') {
                        continue;
                    }
                    item[key] = e.data[key];
                }
            }

            items.push(item);

            let p = e.firstChild;
            while (p) {
                add(p);
                p = p.nextSibling;
            }
        }

        add(this.root);

        return {
            id: this._id,
            items: items
        }
    }

    restore(data: Data): void {
        if (this._queue !== undefined) {
            this._queue.clear();
        }
        this._id = data.id;
        this._elementSet = {};

        for (let item of data.items) {
            let id = item.$id
            let pid = item.$pid
            let data: any = {};
            for (let key in item) {
                if (key.length > 0 && key.substr(0, 1) == '$') {
                    continue;
                }
                data[key] = item[key];
            }
            if (id == 0) {
                this._root = new Element(this, id);
                this._root.data = data;
            } else {
                let e = new Element(this, id);
                e.data = data;
                this._elementSet[id] = e;
                if (pid === 0) {
                    this._root.add(e);
                } else if (pid !== undefined) {
                    let p = this._elementSet[pid];
                    if (p !== undefined) {
                        p.add(e);
                    }
                }
            }
        }

    }

    get(id: Id): Element | undefined {
        return this._elementSet[id];
    }

    create(): Element {
        let id = ++this._id;
        let e = new Element(this, id);
        this._elementSet[id] = e;
        if (this._queue !== undefined) {
            this._queue.add({
                cancel: () => {
                    this._elementSet[id];
                },
                redo: () => {
                    this._elementSet[id] = e;
                }
            });
        }
        return e;
    }

    del(id: Id): void {
        let v = this._elementSet[id];
        if (v !== undefined) {
            delete this._elementSet[id];
            if (this._queue !== undefined) {
                this._queue.add({
                    cancel: () => {
                        this._elementSet[id] = v;
                    },
                    redo: () => {
                        delete this._elementSet[id];
                    }
                });
            }
        }

    }

    addCommand(v: Command): void {
        if (this._queue !== undefined) {
            this._queue.add(v);
        }
    }

}

