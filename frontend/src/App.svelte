<div class="w-full mx-auto text-center md:w-1/2">
    <h1 class="mb-8 text-2xl">Go/Gin sample app</h1>

    <p class="mb-4">Here's a random quote from Taylor Swift for you:</p>
    {#await p}
        <p>Loading…</p>
    {:then quote}
        <p class="p-4 text-lg text-blue-800 bg-gray-100 shadow-md">{quote}</p>
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
    const res = await fetch((URL_PREFIX || '') + '/api/quote')
    const data = await res.json()
    if (!data || !data.quote) {
        throw Error('Response did not contain a quote!')
    }
    return data.quote
}
</script>