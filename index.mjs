import child_process from 'child_process';

const child = child_process.exec('./test');


child.stdout.on('data', (chunk) => {
    const chunkString = chunk.toString('utf8')
    console.log(JSON.parse(chunkString)?.msg)
})
