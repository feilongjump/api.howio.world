<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource as BaseJsonResource;

class JsonResource extends BaseJsonResource
{
    public function __construct($resource)
    {
        parent::__construct($resource);
    }

    public static function collection($resource): \Illuminate\Http\Resources\Json\AnonymousResourceCollection
    {
        return parent::collection($resource);
    }

}
