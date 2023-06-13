<?php

namespace App\Http\Controllers;

use App\Models\Post;
use Illuminate\Http\Request;
use App\Http\Resources\PostResource;
use App\Http\Requests\StorePostRequest;
use App\Http\Requests\UpdatePostRequest;
use Illuminate\Database\Eloquent\Builder;

class PostController extends Controller
{
    public function __construct()
    {
        $this->middleware('auth')->except([
            'index', 'show'
        ]);
    }

    /**
     * Display a listing of the resource.
     */
    public function index(Request $request)
    {
        $posts = Post::with('user')
            ->where(function (Builder $query) use ($request) {
                $query->published()
                    ->when($request->user(),
                        fn (Builder $query, $user) => $query->orWhere('user_id', $user->id)
                    );
            })
            ->when($request->keywords,
                fn (Builder $query, $keywords) => $query->keywords($keywords)
            )
            ->orderByDesc('published_at')
            ->simplePaginate($request->get('per_page', 20));

        return PostResource::collection($posts);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StorePostRequest $request)
    {
        $this->authorize('create', Post::class);

        $request->merge([
            'user_id' => request()->user()->id
        ]);
        $post = Post::create($request->all());

        return new PostResource($post);
    }

    /**
     * Display the specified resource.
     */
    public function show(Post $post)
    {
        $this->authorize('view', $post);

        $post->load(['user', 'content']);

        return new PostResource($post);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdatePostRequest $request, Post $post)
    {
        $this->authorize('own', $post);

        $post->update($request->all());

        return new PostResource($post);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Post $post)
    {
        $this->authorize('own', $post);

        $post->delete();

        return response()->noContent();
    }
}
