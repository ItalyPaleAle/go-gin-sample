<div class="text-center w-full md:w-1/2 mx-auto">
    <h1 class="text-2xl mb-8">Go/Gin sample app</h1>

    <p class="mb-4">Here's a random quote from Taylor Swift for you:</p>
    {#await p}
        <p>Loading…</p>
    {:then quote}
        <p class="p-4 text-lg bg-gray-100 text-blue-800 shadow-md">{quote}</p>
    {:catch err}
        <p class="text-red-800">
            Oopsie, we hit a snag:
            <br />{err}
        </p>
    {/await}
</div>

<script>
let p = requestQuote()

async function requestQuote() {
    const res = await fetch('/api/quote')
    const data = await res.json()
    if (!data || !data.quote) {
        throw Error('Response did not contain a quote!')
    }
    return data.quote
}
</script>