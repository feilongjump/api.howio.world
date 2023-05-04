<?php

namespace App\Http\Controllers;

use App\Models\Thread;
use Illuminate\Http\Request;
use App\Http\Resources\ThreadResource;
use App\Http\Requests\StoreThreadRequest;
use App\Http\Requests\UpdateThreadRequest;

class ThreadController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Request $request)
    {
        $threads = Thread::where('user_id', $request->user()->id)
            ->orderByDesc('published_at')
            ->paginate($request->get('per_page', 20));

        return ThreadResource::collection($threads);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreThreadRequest $request)
    {
        $this->authorize('create', Thread::class);

        $request->merge([
            'user_id' => request()->user()->id
        ]);
        $thread = Thread::create($request->all());

        return new ThreadResource($thread);
    }

    /**
     * Display the specified resource.
     */
    public function show(Thread $thread)
    {
        $this->authorize('own', $thread);

        $thread->load('content');

        return new ThreadResource($thread);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateThreadRequest $request, Thread $thread)
    {
        $this->authorize('own', $thread);

        $thread->update($request->all());

        return new ThreadResource($thread);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Thread $thread)
    {
        $this->authorize('own', $thread);

        $thread->delete();

        return response()->noContent();
    }
}
